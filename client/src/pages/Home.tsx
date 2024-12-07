import { SignInForm } from "@/components/signin-form";
import { SignoutButton } from "@/components/signout-button";
import { useAuth } from "@/hooks/use-auth";

export default function Home() {
  const { isAuthenticated, email, isLoading } = useAuth();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return isAuthenticated ? (
    <div>
      <h1>Welcome to the Lake Braddock Ultimate Club Admin Dashboard</h1>
      <p>Logged in as: {email}</p>
      <SignoutButton />
    </div>
  ) : (
    <div className="flex justify-center items-center h-screen w-full">
      <SignInForm />
    </div>
  );
}