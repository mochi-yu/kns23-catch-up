"use client";

import { auth } from "@/lib/firebase/firebase";
import { Button } from "@mui/material";
import { GoogleAuthProvider, signInWithPopup } from "firebase/auth";

export function GoogleLoginButton() {
  // ユーザーがログインボタンを押したときにdoLogin関数が実行される
  const loginByGoogle = () => {
    // Firebaseで用意されているメールアドレスとパスワードでログインするための関数
    const googleProvider = new GoogleAuthProvider();
    signInWithPopup(auth, googleProvider).catch((error) => {
      console.log(error);
    });
  };

  return (
    <Button variant='contained' onClick={loginByGoogle}>
      Googleでログイン
    </Button>
  );
}
