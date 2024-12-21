import { Player } from "@/types"
import { CheckTable } from "@/components/check-table";
import { API_BASE_URL } from "@/config";
import { Team, Check } from "@/types";
import { ColumnDef } from "@tanstack/react-table";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom"
import { Checkbox } from "@/components/ui/checkbox";

export default function Checklist() {
  const { id } = useParams()

  const [team, setTeam] = useState<Team>();
  const [checks, setChecks] = useState<Record<number, Check[]>>({});

  useEffect(() => {
    const fetchTeam = async () => {
      const res = await fetch(`${API_BASE_URL}/team/${id}`)
      const data = await res.json() as Team
      setTeam(data)
    }

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

  const columns: ColumnDef<Player>[] = [
    { header: "Player Name", accessorKey: "player_name" },
  ]

  if (team?.tasks) {
    for (const task of team.tasks) {
      columns.push({
        header: task.description,
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
      <div className="container mx-auto py-10">
        <CheckTable columns={columns} data={team?.players ?? []} checks={checks} />
      </div>
    </div>

  )
}