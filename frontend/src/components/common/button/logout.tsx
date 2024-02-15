"use client";

import { auth } from "@/lib/firebase/firebase";
import { Button } from "@mui/material";
import { signOut } from "firebase/auth";

export function LogoutButton() {
  const logout = () => {
    signOut(auth);
  };

  return (
    <Button variant='contained' onClick={logout}>
      ログアウト
    </Button>
  );
}
