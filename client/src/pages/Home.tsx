import { SignInForm } from "@/components/signin-form";
import { useAuth } from "@/hooks/use-auth";

export default function Home() {
  const { isAuthenticated, isLoading } = useAuth();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return isAuthenticated ? (
    <div>
      <h1>Welcome to the Lake Braddock Ultimate Club Admin Dashboard</h1>
    </div>
  ) : (
    <div className="flex justify-center items-center h-screen w-full">
      <SignInForm />
    </div>
  );
} 