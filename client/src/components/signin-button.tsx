import { Button } from "@/components/ui/button";
import { useAuth } from "@/hooks/use-auth";

export function SigninButton() {
    const { signin } = useAuth();

    return (
        <>
            <Button onClick={signin}>Sign in with Google</Button>
        </>
    );
}       