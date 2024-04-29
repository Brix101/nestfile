import { Button } from "@/components/ui/button";
import { QUERY_KEYS } from "@/constant/query-key";
import { useQueryClient } from "@tanstack/react-query";

function FileListingPage() {
  const queryClient = useQueryClient();

  return (
    <div>
      File Listing Page
      <Button
        onClick={() =>
          queryClient.setQueryData([QUERY_KEYS.auth_user], { user: null })
        }
      >
        Logout
      </Button>
    </div>
  );
}

export default FileListingPage;
