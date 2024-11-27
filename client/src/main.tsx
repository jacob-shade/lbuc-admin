import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar"
import { AppSidebar } from "@/components/app-sidebar"

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <SidebarProvider>
      <div className="flex min-h-screen">
        <div className="flex-shrink-0">
          <AppSidebar />
        </div>
        <div className="flex flex-1 flex-col">
          <div className="flex items-center p-4">
            <SidebarTrigger />
          </div>
          <main>
            <App />
          </main>
        </div>
      </div>
    </SidebarProvider>
  </StrictMode> 
)
