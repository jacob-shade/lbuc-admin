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
            check => check.player_id === playerId && check.checked
          );

          return (
            <Checkbox
              checked={isChecked}
              onCheckedChange={(checked) => {
                // Here you would add API call to update the check status
                console.log(`Player ${playerId}, Task ${taskId}, Checked: ${checked}`);
              }}
            />
          );
        }
      });
    }
  }

  console.log("checks", checks)

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">{team?.team_name} Checklist</h1>
      <div className="container mx-auto py-10">
        <CheckTable columns={columns} data={team?.players ?? []} checks={checks} />
      </div>
    </div>

  )
}