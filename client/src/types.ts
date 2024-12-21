interface Team {
    id: number;
    team_name: string;
    players: Player[];
    tasks: Task[];
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

interface Task {
    id: number;
    description: string;
}

interface Check {
    player_id: number;
    task_id: number;
    checked: boolean;
}

export type { Team, Player, Task, Check }
