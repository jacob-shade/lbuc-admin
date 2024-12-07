import { useAuth } from "@/hooks/use-auth";
import webNeutralSvg from "@/assets/web_neutral_rd_SI.svg";

export function SigninButton() {
    const { signin } = useAuth();

    return (
        <button
            type="button"
            className="w-full flex justify-center"
            onClick={signin}
        >
            <img src={webNeutralSvg} alt="Google sign-in" className="h-10 w-auto" />
        </button>
    )
}       