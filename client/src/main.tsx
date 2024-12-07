import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { SidebarProvider } from "@/components/ui/sidebar"
import { GoogleOAuthProvider } from '@react-oauth/google'
import { AuthProvider } from './context/auth-context-provider.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <GoogleOAuthProvider clientId={import.meta.env.VITE_GOOGLE_CLIENT_ID}>
      <SidebarProvider>
        <AuthProvider>
          <App />
        </AuthProvider>
      </SidebarProvider>
    </GoogleOAuthProvider>
  </StrictMode>
)