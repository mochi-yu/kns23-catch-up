"use client";

import { AuthContext } from "@/contexts/auth";
import { redirect } from "next/navigation";
import { useContext, useEffect } from "react";

interface Props {
  children: React.ReactNode;
}

export default function CheckNotRegister({ children }: Props) {
  const { currentUser, userInfo } = useContext(AuthContext);

  useEffect(() => {
    if (!currentUser) {
      // ログインしていなかったら
      redirect("/");
    } else if (userInfo?.is_temp_user == false) {
      // tempユーザではなかったら
      redirect("/home");
    }
  });

  return <>{children}</>;
}
