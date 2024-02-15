"use client";

import { AuthContext } from "@/contexts/auth";
import { redirect } from "next/navigation";
import { ReactNode, useContext, useEffect } from "react";

interface Props {
  children: ReactNode;
}

export default function CheckLoginUser({ children }: Props) {
  const { currentUser, userInfo } = useContext(AuthContext);

  useEffect(() => {
    if (!currentUser) {
      // Firebaseでログインできていなかったら
      redirect("/");
    } else if (userInfo?.is_temp_user == true) {
      // 正規ユーザではなかったら
      redirect("/register");
    }
  });

  return <>{children}</>;
}
