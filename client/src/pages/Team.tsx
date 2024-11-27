import { useState, useEffect } from "react";
import { useParams } from "react-router-dom"

export default function Team() {
  const { id } = useParams()

  const BASE_URL = "http://localhost:5000"

  interface Team {
      id: number;
      team_name: string;
      players: Player[];
  }

  interface Player {
      id: number;
      player_name: string;
  }

  const [team, setTeam] = useState<Team>();

  useEffect(() => {
      const fetchTeam = async () => {
          const res = await fetch(`${BASE_URL}/api/team/${id}`)
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