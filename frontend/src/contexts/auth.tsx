"use client";

import { createContext, useEffect, useState, ReactNode } from "react";
import { auth } from "@/lib/firebase/firebase";
import { User } from "firebase/auth";
import { Loading } from "@/components/common/loading";
import { UserModel } from "@/model/user";
import axios from "@/util/axios";
import { AxiosResponse } from "axios";

type AuthContextProps = {
  currentUser: User | null | undefined;
  userInfo: UserModel | undefined;
  signInCheck: boolean;
};

const AuthContext = createContext<AuthContextProps>({
  currentUser: undefined,
  userInfo: undefined,
  signInCheck: false,
});

interface Props {
  children: ReactNode;
}

function AuthProvider({ children }: Props) {
  const [currentUser, setCurrentUser] = useState<User | null | undefined>(undefined);
  const [signInCheck, setSignInCheck] = useState(false);
  const [userInfo, setUserInfo] = useState<UserModel | undefined>(undefined);

  // ログイン状態を確認する
  useEffect(() => {
    auth.onAuthStateChanged(async (user) => {
      if (user) {
        // ログイン状態
        setCurrentUser(user);

        // サーバからユーザ情報を取得する
        const idToken = await currentUser?.getIdToken();
        if (!idToken) return;

        const userInfo = await axios
          .get("/auth", {
            headers: { Authorization: "Bearer " + idToken },
          })
          .then((res: AxiosResponse<UserModel>) => {
            const { data } = res;
            return data;
          })
          .catch((e) => {
            console.log(e);
          });
        setUserInfo(userInfo!);

        setSignInCheck(true);
      } else {
        // ログアウトの状態では、各stateをundefineにする
        setCurrentUser(undefined);
        setUserInfo(undefined);
        setSignInCheck(true);
      }
    });
  });

  if (signInCheck) {
    return (
      <AuthContext.Provider value={{ currentUser, userInfo, signInCheck }}>
        {children}
      </AuthContext.Provider>
    );
  } else {
    // ログイン確認中
    // 自分で作ったローディングコンポーネントをレンダリングする
    return <Loading open={true} />;
  }
}

export { AuthContext, AuthProvider };
