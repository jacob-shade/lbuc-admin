interface Team {
    id: number;
    team_name: string;
    players: Player[];
}

interface Player {
    id: number;
    player_name: string;
    nick_name: string;
    pronouns: string;
    grade: number;
    birthday: Date;
    player_email: string;
    parent_name: string;
    parent_email: string;
    parent_number: string;
    relationship: string;
    address: string;
    medical_notes: string;
}

export type { Team, Player }
