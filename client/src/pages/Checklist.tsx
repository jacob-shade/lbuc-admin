import { Player } from "@/types"
import { CheckTable } from "@/components/check-table";
import { BASE_URL } from "@/config";
import { Team } from "@/types";
import { ColumnDef } from "@tanstack/react-table";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom"

export default function Checklist() {
  const { id } = useParams()

  const [team, setTeam] = useState<Team>();

  useEffect(() => {
    const fetchTeam = async () => {
      const res = await fetch(`${BASE_URL}/api/team/${id}`)
      const data = await res.json() as Team
      setTeam(data)
    }

    fetchTeam()
  }, [id])

  const columns: ColumnDef<Player>[] = [
    { header: "Player Name", accessorKey: "player_name" },
    { header: "LBUC Registration", accessorKey: "lbuc_reg" },
    { header: "USAU Registration", accessorKey: "usau_reg" },
  ]

  //console.log("EC team",team)
  //console.log("EC players",team?.players)
  //console.log("EC players[0]",team?.players[0])

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">{team?.team_name} Checklist</h1>
      <div className="container mx-auto py-10">
        <CheckTable columns={columns} data={team?.players ?? []} />
      </div>
    </div>

  )
}