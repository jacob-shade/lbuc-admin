interface Team {
    id: number;
    team_name: string;
    players: Player[];
}

interface Player {
    id: number;
    player_name: string;
}

export type { Team, Player }
