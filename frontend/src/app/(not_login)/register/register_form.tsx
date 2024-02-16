"use client";

import { AuthContext } from "@/contexts/auth";
import { Button, FormHelperText, MenuItem, Stack, TextField } from "@mui/material";
import { useContext, useEffect, useState } from "react";

export const ClassList = ["23A", "23B", "23C", "23D", "23E", "23F", "23G", "23H", "23J"];

export function RegisterForm() {
  const { currentUser } = useContext(AuthContext);

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
    console.log(e);
    setUserID(e.target.value);
  }
  function onDisplayNameChange(e: any) {
    console.log(e);
    setDisplayName(e.target.value);
  }
  function onUserNameChange(e: any) {
    console.log(e);
    setUserName(e.target.value);
  }
  function onClassIDChange(e: any) {
    console.log(e);
    setClassID(e.target.value);
  }

  function onSubmit() {
    if (userID == "" || displayName == "" || userName == "" || classID == "") {
      console.log("invalid param.");
      return;
    } else {
      console.log("submit.");
    }
  }

  return (
    <>
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
