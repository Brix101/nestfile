import { useUser } from "@/hooks/useUser";
import { Navigate } from "react-router-dom";

function LoginPage() {
  const { user, isLoaded } = useUser();

  if (isLoaded && user) {
    return <Navigate to="/files" />;
  }

  return <div>Login Page</div>;
}

export default LoginPage;
