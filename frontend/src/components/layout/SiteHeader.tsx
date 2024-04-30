import { CommandMenu } from "@/components/CommandMenu";
import ModeToggle from "@/components/ModeToggle";
import { Icons } from "@/components/icons";
import { siteConfig } from "@/config/site";

function SiteHeader() {
  return (
    <header className="sticky top-0 z-50 w-full border-b border-border/40 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 shadow-md">
      <div className="container flex justify-between h-14 max-w-screen-2xl items-center">
        <div className="flex gap-6 md:gap-10">
          <div className="flex items-center space-x-2">
            <Icons.logo className="h-6 w-6" />
            <span className="inline-block font-bold">{siteConfig.name}</span>
          </div>
        </div>
        <div className="md:w-auto md:flex-none">
          <CommandMenu />
        </div>
        <div className="flex items-center justify-between space-x-2 md:justify-end">
          <nav className="flex items-center">
            <ModeToggle />
          </nav>
        </div>
      </div>
    </header>
  );
}

export default SiteHeader;
