import { SidebarTrigger } from "@/components/ui/sidebar"
import { AppSidebar } from "@/components/app-sidebar"

export function Layout({ children }: { children: React.ReactNode }) {
  return (
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
  )
}