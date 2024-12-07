import { Player } from "@/types"
import { DataTable } from "@/components/data-table";
import { API_BASE_URL } from "@/config";
import { Team } from "@/types";
import { ColumnDef } from "@tanstack/react-table";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom"

export default function EmergencyContacts() {
  const { id } = useParams()

  const [team, setTeam] = useState<Team>();

  useEffect(() => {
    const fetchTeam = async () => {
      const res = await fetch(`${API_BASE_URL}/team/${id}`)
      const data = await res.json() as Team
      setTeam(data)
    }

    fetchTeam()
  }, [id])

  const columns: ColumnDef<Player>[] = [
    { header: "Player Name", accessorKey: "player_name" },
    { header: "Nickname", accessorKey: "nick_name" },
    { header: "Pronouns", accessorKey: "pronouns" },
    { header: "Grade", accessorKey: "grade" },
    {
      header: "Birthday",
      accessorKey: "birthday",
      cell: ({ row }) => {
        const date = new Date(row.getValue("birthday"))
        return date.toLocaleDateString("en-US", {
          month: "2-digit",
          day: "2-digit",
          year: "numeric"
        }).replace(/\//g, "-")
      }
    },
    { header: "Player Email", accessorKey: "player_email" },
    { header: "Parent Name", accessorKey: "parent_name" },
    { header: "Parent Email", accessorKey: "parent_email" },
    { header: "Parent Number", accessorKey: "parent_number" },
    { header: "Relationship", accessorKey: "relationship" },
    { header: "Address", accessorKey: "address" },
    { header: "Medical Notes", accessorKey: "medical_notes" },
  ]

  //   console.log("EC team",team)
  //   console.log("EC players",team?.players)
  //   console.log("EC players[0]",team?.players[0])

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">{team?.team_name} Emergency Contacts</h1>
      <div className="container mx-auto py-10">
        <DataTable columns={columns} data={team?.players ?? []} />
      </div>
    </div>

  )
}