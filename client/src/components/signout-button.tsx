import { Button } from "./ui/button";
import { useAuth } from "@/hooks/use-auth";

export function SignoutButton() {
    const { signout } = useAuth();

    return (
        <Button onClick={signout} variant="outline">
            Sign Out
        </Button>
    );
}