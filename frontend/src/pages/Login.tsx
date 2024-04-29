import { UserAuthForm } from "@/components/forms/user-auth-form";
import { Card } from "@/components/ui/card";

function LoginPage() {
  return (
    <>
      <div className="container relative flex h-screen  items-center justify-center ">
        <Card className="lg:p-8">
          <div className="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
            <div className="flex flex-col space-y-2 text-center">
              <h1 className="text-2xl font-semibold tracking-tight">Login</h1>
              <p className="text-sm text-muted-foreground">
                Enter your account
              </p>
            </div>
            <UserAuthForm />
          </div>
        </Card>
      </div>
    </>
  );
}

export default LoginPage;
