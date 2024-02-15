import CheckRootPage from "@/components/auth/check_root_page";
import { LogoutButton } from "@/components/common/button/logout";
import { GoogleLoginButton } from "@/components/common/button/login_google";

export default function RootPage() {
  return (
    <>
      <CheckRootPage>
        <GoogleLoginButton />
        <LogoutButton />
      </CheckRootPage>
    </>
  );
}
