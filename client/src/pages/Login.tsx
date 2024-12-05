import { useGoogleLogin } from "@react-oauth/google"

export default function Login() {
    const login = useGoogleLogin({
        onSuccess: codeResponse => console.log(codeResponse),
        flow: 'auth-code',
    });

    return <button onClick={login}>Login</button>
}


