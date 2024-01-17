"use client";

import { Stack, IconButton, InputBase, styled } from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";
import { SyntheticEvent, useState } from "react";

const StyledStack = styled(Stack)({
  backgroundColor: "#EDEDF4",
  borderRadius: "15px",
  padding: "10px 10px 10px 15px",
  height: "100%",
});

export function SearchBox() {
  const [searchText, setSearchText] = useState<string>("");

  function searchButtonClickHandler() {
    console.log("pushed search button.");
  }

  return (
    <>
      <StyledStack direction='row'>
        <InputBase
          sx={{ ml: 1, flex: 1 }}
          placeholder='ここに検索する文字列を入力'
          value={searchText}
          onChange={(e) => {
            setSearchText(e.target.value);
          }}
        />
        <IconButton
          type='button'
          sx={{ m: "-10px 0px -10px 0px" }}
          onClick={searchButtonClickHandler}
        >
          <SearchIcon />
        </IconButton>
      </StyledStack>
    </>
  );
}
