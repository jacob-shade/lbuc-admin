import { SidebarTrigger } from "@/components/ui/sidebar"
import { AppSidebar } from "@/components/app-sidebar"
import { useAuth } from "@/hooks/use-auth";

export function Layout({ children }: { children: React.ReactNode }) {
  const { isAuthenticated } = useAuth();

  return isAuthenticated ? (
    <div className="flex">
      <AppSidebar />
      <main className="flex">
        <div className="p-3">
          <SidebarTrigger />
        </div>
        <div className="p-3">
          {children}
        </div>
      </main>
    </div>
  ) : (
    <div className="w-full">
      {children}
    </div>
  );
}