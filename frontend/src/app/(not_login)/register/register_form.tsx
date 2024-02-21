"use client";

import { AuthContext, SetUserInfoContext } from "@/contexts/auth";
import { Button, FormHelperText, MenuItem, Stack, TextField } from "@mui/material";
import { useContext, useEffect, useState } from "react";
import { PostWithLogin } from "@/util/axios";
import { UserModel } from "@/model/user";
import { Loading } from "@/components/common/loading";
import { useRouter } from "next/navigation";

export const ClassList = ["23A", "23B", "23C", "23D", "23E", "23F", "23G", "23H", "23J"];

export function RegisterForm() {
  const { currentUser } = useContext(AuthContext);
  const { setUserInfo } = useContext(SetUserInfoContext);
  const router = useRouter();

  const [isLoading, setIsLoading] = useState(false);
  const [userID, setUserID] = useState("");
  const [displayName, setDisplayName] = useState("");
  const [userName, setUserName] = useState("");
  const [classID, setClassID] = useState("");
  const [mailAddress, setMailAddress] = useState("");

  useEffect(() => {
    if (currentUser?.email) {
      setMailAddress(currentUser.email);
    }
  }, [currentUser]);

  function onUserIDChange(e: any) {
    setUserID(e.target.value);
  }
  function onDisplayNameChange(e: any) {
    setDisplayName(e.target.value);
  }
  function onUserNameChange(e: any) {
    setUserName(e.target.value);
  }
  function onClassIDChange(e: any) {
    setClassID(e.target.value);
  }

  async function onSubmit() {
    if (userID == "" || displayName == "" || userName == "" || classID == "") {
      console.log("invalid param.");
      return;
    } else {
      setIsLoading(true);

      const reqParam = {
        user_id: userID,
        display_name: displayName,
        user_name: userName,
        class_id: classID,
      };

      const userInfo = (await PostWithLogin("/auth", reqParam)) as UserModel;
      setUserInfo(userInfo);
      setIsLoading(false);
    }
  }

  return (
    <>
      <Loading open={isLoading} />
      <Stack margin='20px' spacing={2}>
        <TextField
          variant='outlined'
          label='本名（本人確認に使用します）'
          value={userName}
          onChange={onUserNameChange}
        />
        <FormHelperText>例）神奈総 太郎</FormHelperText>

        <TextField
          variant='outlined'
          label='ユーザID（半角英数）'
          value={userID}
          onChange={onUserIDChange}
        />
        <FormHelperText>例）kns-taro</FormHelperText>

        <TextField
          variant='outlined'
          label='表示名'
          value={displayName}
          onChange={onDisplayNameChange}
        />
        <FormHelperText>例）かなそー生</FormHelperText>

        <TextField variant='outlined' label='メールアドレス' value={mailAddress} disabled />
        <TextField
          select
          label='神奈総のクラス'
          variant='outlined'
          value={classID}
          onChange={onClassIDChange}
        >
          {ClassList.map((elm, i) => (
            <MenuItem key={elm} value={elm}>
              {elm}
            </MenuItem>
          ))}
        </TextField>

        <Button variant='contained' onClick={onSubmit}>
          登録
        </Button>
      </Stack>
    </>
  );
}
