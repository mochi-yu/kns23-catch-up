"use client";

import { auth } from "@/lib/firebase/firebase";
import { Button } from "@mui/material";
import { GoogleAuthProvider, signInWithPopup } from "firebase/auth";

export function GoogleLoginButton() {
  // ユーザーがログインボタンを押したときにdoLogin関数が実行される
  const loginByGoogle = () => {
    // Firebaseで用意されているメールアドレスとパスワードでログインするための関数
    const googleProvider = new GoogleAuthProvider();
    signInWithPopup(auth, googleProvider)
      .then((result) => {
        // This gives you a Google Access Token. You can use it to access the Google API.
        const credential = GoogleAuthProvider.credentialFromResult(result);
        const token = credential?.accessToken;
        console.log(credential);
        // The signed-in user info.
        const user = result.user;
        console.log(result);
        auth.currentUser?.getIdToken().then((res) => {
          console.log(res);
        });
        // IdP data available using getAdditionalUserInfo(result)
        // ...
      })
      .catch((error) => {
        // Handle Errors here.
        const errorCode = error.code;
        const errorMessage = error.message;
        // The email of the user's account used.
        const email = error.customData.email;
        // The AuthCredential type that was used.
        const credential = GoogleAuthProvider.credentialFromError(error);
        // ...
      });
  };

  return (
    <Button variant='contained' onClick={loginByGoogle}>
      Googleでログイン
    </Button>
  );
}
