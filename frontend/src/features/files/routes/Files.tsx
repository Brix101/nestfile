import { Icons } from "@/components/icons";
import { Button } from "@/components/ui/button";
import { useLogout } from "@/features/auth/api/logout";

export function Files() {
  const { mutate, isPending } = useLogout();

  return (
    <div>
      File Listing Page
      <Button onClick={() => mutate()}>
        {isPending && (
          <Icons.spinner
            className="mr-2 size-4 animate-spin"
            aria-hidden="true"
          />
        )}
        Logout
        <span className="sr-only">Logout</span>
      </Button>
    </div>
  );
}
