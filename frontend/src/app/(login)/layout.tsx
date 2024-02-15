import { Header } from "@/components/common/header/header";
import { Footer } from "@/components/common/footer";
import CheckLoginUser from "@/components/auth/check_login_user";

interface Props {
  children: React.ReactNode;
}

export default function LoginLayout({ children }: Props) {
  return (
    <>
      <CheckLoginUser>
        <Header />
        {children}
        <Footer />
      </CheckLoginUser>
    </>
  );
}
