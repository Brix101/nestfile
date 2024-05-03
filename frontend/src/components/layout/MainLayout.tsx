import React from "react";

import SiteHeader from "./SiteHeader";
import { useUser } from "@/hooks/useUser";
import { SiteSideBar } from "./SiteSidebar";

interface MainLayoutProps extends React.PropsWithChildren {}

export function MainLayout({ children }: MainLayoutProps) {
  const { user } = useUser();

  return (
    <div className="flex min-h-screen flex-col">
      <SiteHeader user={user} />
      <div className="container flex-1 items-start lg:grid lg:grid-cols-[240px_minmax(0,1fr)] lg:gap-10">
        <SiteSideBar
          // the top-16 class is used for the dashboard-header of h-16, added extra 0.1rem to fix the sticky layout shift issue
          className="top-[calc(theme('spacing.16')_+_0.1rem)] z-30 hidden border-r lg:sticky lg:block"
          user={user}
        />
        <main className="flex w-full flex-col overflow-hidden">{children}</main>
      </div>
    </div>
  );
}
