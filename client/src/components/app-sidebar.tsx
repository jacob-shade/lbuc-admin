import * as React from "react"
import { useState, useEffect } from "react"
import { GalleryVerticalEnd } from "lucide-react"

import {
  Sidebar,
  SidebarContent,
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

const teamLogistics = ["Checklist", "Emergency Contacts", "Email List"]

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
    const BASE_URL = "http://localhost:5000"

    interface Team {
        id: number;
        team_name: string;
    }

    const [teams, setTeams] = useState<Team[]>([]);

    useEffect(() => {
        const fetchTeams = async () => {
            const res = await fetch(`${BASE_URL}/api/team`)
            const data = await res.json() as Team[]
            setTeams(data)
        }
        fetchTeams()
    }, [])

    console.log("TEAMS", teams)

  return (
    <Sidebar {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" asChild>
              <a href="#">
                <div className="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                  <GalleryVerticalEnd className="size-4" />
                </div>
                <div className="flex flex-col gap-0.5 leading-none">
                  <span className="font-semibold">Lake Braddock Ultimate Club</span>
                  <span className="">Admin Dashboard</span>
                </div>
              </a>
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
                  <a href={`/team/${team.id}`} className="font-medium">
                    {team.team_name}
                  </a>
                </SidebarMenuButton>
                {teamLogistics.length ? (
                  <SidebarMenuSub>
                    {teamLogistics.map((item) => (
                      <SidebarMenuSubItem key={item}>
                        <SidebarMenuSubButton asChild>
                          <a href="#">{item}</a>
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
      <SidebarRail />
    </Sidebar>
  )
}
