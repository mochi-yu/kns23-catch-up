import { Header } from "@/components/common/header/header";
import { Footer } from "@/components/common/footer";
import React from "react";

interface Props {
  children: React.ReactNode;
}

export default function LoginLayout({ children }: Props) {
  return (
    <>
      <Header />
      {children}
      <Footer />
    </>
  );
}
