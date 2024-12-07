import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { SigninButton } from "./google-signin-button"
import LBUCLogoSvg from "@/assets/disclbnoclaw.png"

export const description =
    "Sign in with Google to access the LBUC Admin Dashboard."

export function SignInForm() {
    return (
        <Card className='mx-auto max-w-sm'>
            <CardHeader>
                <img src={LBUCLogoSvg} alt="LBUC Logo" className="h-150 w-auto" />
                <CardTitle className='text-2xl text-center'>LBUC Admin Sign In</CardTitle>
                <CardDescription className='text-center'>
                    Must be an approved admin to sign in.
                </CardDescription>
            </CardHeader>
            <CardContent>
                <SigninButton />
            </CardContent>
        </Card>
    )
}