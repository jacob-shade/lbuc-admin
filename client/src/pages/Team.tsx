import { useState, useEffect } from "react";
import { useParams } from "react-router-dom"
import type { Team } from "@/types";
import { API_BASE_URL } from "@/config";

export default function Team() {
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

  return (
    <div>
      <h1>Team {team?.team_name}</h1>
      <div>Players</div>
      {team?.players.map((player) => (
        <div key={player.id}>{player.player_name}</div>
      ))}
    </div>
  )
}