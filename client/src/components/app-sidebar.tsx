import * as React from "react"
import { useState, useEffect } from "react"
import { GalleryVerticalEnd } from "lucide-react"
import { Link } from "react-router-dom"
import { API_BASE_URL } from "@/config"
import { Team } from "@/types"
import { NavUser } from "@/components/nav-user"
import { useAuth } from "@/hooks/use-auth"

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
  SidebarRail,
} from "@/components/ui/sidebar"

const teamLogistics = [
  { name: "Checklist", path: "/checklist" },
  { name: "Emergency Contacts", path: "/emergency-contacts" },
  { name: "Email List", path: "/email-list" },
]

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const [teams, setTeams] = useState<Team[]>([]);
  const { email, name, picture } = useAuth();

  useEffect(() => {
    const fetchTeams = async () => {
      const res = await fetch(`${API_BASE_URL}/team`)
      const data = await res.json() as Team[]
      setTeams(data)
    }
    fetchTeams()
  }, [])

  return (
    <Sidebar {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" asChild>
              <Link to="/">
                <div className="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                  <GalleryVerticalEnd className="size-4" />
                </div>
                <div className="flex flex-col gap-0.5 leading-none">
                  <span className="font-semibold">Lake Braddock Ultimate Club</span>
                  <span className="">Admin Dashboard</span>
                </div>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarMenu>
            {teams.map((team) => (
              <SidebarMenuItem key={team.id}>
                <SidebarMenuButton asChild>
                  <Link to={`/team/${team.id}`} className="font-medium">
                    {team.team_name}
                  </Link>
                </SidebarMenuButton>
                {teamLogistics.length ? (
                  <SidebarMenuSub>
                    {teamLogistics.map((item) => (
                      <SidebarMenuSubItem key={item.name}>
                        <SidebarMenuSubButton asChild>
                          <Link to={`${item.path}/${team.id}`}>{item.name}</Link>
                        </SidebarMenuSubButton>
                      </SidebarMenuSubItem>
                    ))}
                  </SidebarMenuSub>
                ) : null}
              </SidebarMenuItem>
            ))}
          </SidebarMenu>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={{ name: name ?? "", email: email ?? "", avatar: picture ?? "" }} />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  )
}
