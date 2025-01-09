import { Player } from "@/types"
import { CheckTable } from "@/components/check-table";
import { API_BASE_URL } from "@/config";
import { Team, Check } from "@/types";
import { ColumnDef } from "@tanstack/react-table";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom"
import { Checkbox } from "@/components/ui/checkbox";
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu"
import { MoreHorizontal } from "lucide-react"

export default function Checklist() {
  const { id } = useParams()

  const [team, setTeam] = useState<Team>();
  const [checks, setChecks] = useState<Record<number, Check[]>>({});
  const [newTaskDescription, setNewTaskDescription] = useState<string>("");
  const [open, setOpen] = useState(false);
  const [taskToDelete, setTaskToDelete] = useState<number | null>(null);
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);

  const handleDeleteClick = (taskId: number) => {
    setTaskToDelete(taskId);
    setDeleteDialogOpen(true);
  };

  const fetchTeam = async () => {
    const res = await fetch(`${API_BASE_URL}/team/${id}`)
    const data = await res.json() as Team
    setTeam(data)
  }

  useEffect(() => {
    fetchTeam()
  }, [id])

  useEffect(() => {
    const fetchAllChecks = async () => {
      if (!team?.tasks) return;

      const checksByTask: Record<number, Check[]> = {};

      for (const task of team.tasks) {
        const res = await fetch(`${API_BASE_URL}/task/${task.id}/checks`);
        const data = await res.json() as Check[];
        checksByTask[task.id] = data;
      }

      setChecks(checksByTask);
    };

    fetchAllChecks();
  }, [team]);

  const createTask = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!newTaskDescription.trim() || !team) return;

    try {
      const response = await fetch(`${API_BASE_URL}/team/${id}/task`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          team_id: team.id,
          description: newTaskDescription
        })
      });

      if (response.ok) {
        const newTask = await response.json();
        // Update the team state with the new task
        setTeam(prev => prev ? {
          ...prev,
          tasks: [...(prev.tasks || []), newTask]
        } : prev);
        // Initialize empty checks for the new task
        setChecks(prev => ({
          ...prev,
          [newTask.id]: []
        }));
        setNewTaskDescription(""); // Reset input
        setOpen(false);
        await fetchTeam();
      }
    } catch (error) {
      console.error("Failed to create task:", error);
    }
  };

  const confirmDelete = async () => {
    if (!taskToDelete || !team) return;

    try {
      const response = await fetch(`${API_BASE_URL}/team/${team.id}/task/${taskToDelete}`, {
        method: "DELETE",
      });

      if (response.ok) {
        setChecks(prev => {
          const newChecks = { ...prev };
          delete newChecks[taskToDelete];
          return newChecks;
        });

        await fetchTeam();
      }
    } catch (error) {
      console.error("Failed to delete task:", error);
    } finally {
      setDeleteDialogOpen(false);
      setTaskToDelete(null);
    }
  };

  const columns: ColumnDef<Player>[] = [
    { id: "player_name", header: "Player Name", accessorKey: "player_name" },
  ]

  if (team?.tasks) {
    for (const task of team.tasks) {
      columns.push({
        id: `task-${task.id}`,
        header: () => (
          <div className="flex items-center justify-between">
            <span className="mr-2">{task.description}</span>
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <Button variant="ghost" className="h-8 w-8 p-0">
                  <MoreHorizontal className="h-4 w-4" />
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent align="end">
                <DropdownMenuItem
                  className="text-red-600"
                  onClick={() => handleDeleteClick(task.id)}
                >
                  Delete Task
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        ),
        cell: ({ row }) => {
          const playerId = row.original.id;
          const taskId = task.id;
          const isChecked = checks[taskId]?.some(
            check => check.PlayerID === playerId && check.Checked
          );

          return (
            <div className="h-6 flex items-center justify-left">
              <Checkbox
                checked={isChecked}
                onCheckedChange={async (checked) => {
                  const check: Check = {
                    TaskID: taskId,
                    PlayerID: playerId,
                    Checked: checked ? true : false
                  }
                  try {
                    const response = await fetch(`${API_BASE_URL}/check`, {
                      method: "POST",
                      headers: {
                        "Content-Type": "application/json"
                      },
                      body: JSON.stringify(check)
                    });
                    if (response.ok) {
                      // Update local state after successful API call
                      setChecks(prevChecks => ({
                        ...prevChecks,
                        [taskId]: [
                          ...(prevChecks[taskId]?.filter(c => c.PlayerID !== playerId) || []),
                          check
                        ]
                      }));
                    }
                  } catch (error) {
                    console.error("Failed to update check:", error);
                    // Optionally add error handling UI here
                  }
                }}
              />
            </div>

          );
        }
      });
    }
  }

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">{team?.team_name} Checklist</h1>
      <Dialog open={open} onOpenChange={setOpen}>
        <DialogTrigger asChild>
          <Button>Add Task</Button>
        </DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Add New Task</DialogTitle>
          </DialogHeader>
          <form onSubmit={createTask} className="grid gap-4 py-4">
            <Input
              placeholder="Task description"
              value={newTaskDescription}
              onChange={(e) => setNewTaskDescription(e.target.value)}
            />
            <Button type="submit">Create Task</Button>
          </form>
        </DialogContent>
      </Dialog>
      <Dialog open={deleteDialogOpen} onOpenChange={setDeleteDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Confirm Delete</DialogTitle>
          </DialogHeader>
          <div className="py-4">
            <p>Are you sure you want to delete this task? This action cannot be undone.</p>
          </div>
          <div className="flex justify-end gap-4">
            <Button
              variant="outline"
              onClick={() => setDeleteDialogOpen(false)}
            >
              Cancel
            </Button>
            <Button
              variant="destructive"
              onClick={confirmDelete}
            >
              Delete
            </Button>
          </div>
        </DialogContent>
      </Dialog>
      <div className="container mx-auto py-10">
        <CheckTable columns={columns} data={team?.players ?? []} checks={checks} />
      </div>
    </div>
  )
}