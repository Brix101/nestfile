import { NewFileMenu } from "@/components/new-file-menu";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { UserResource } from "@/features/auth";
import { cn } from "@/lib/utils";

interface SiteSidebarProps extends React.HTMLAttributes<HTMLElement> {
  children?: React.ReactNode;
  user?: UserResource | null;
}

export function SiteSideBar({
  user,
  children,
  className,
  ...props
}: SiteSidebarProps) {
  return (
    <aside className={cn("w-full", className)} {...props}>
      <ScrollArea className="h-[calc(100vh-8rem)] py-2.5 pr-6">
        <div className="flex flex-col gap-4">
          <NewFileMenu />
          <Separator />
          <Button>Home</Button>
          <Button>My Files</Button>
          <Separator />
          <Button>Settings</Button>
        </div>
      </ScrollArea>
      <div className="pr-6 pt-4 lg:pt-6">{children}</div>
    </aside>
  );
}
