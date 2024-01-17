"use client";

import { Button, Stack, ThemeProvider, createTheme, styled } from "@mui/material";
import EditOutlinedIcon from "@mui/icons-material/EditOutlined";
import Typography from "@mui/material/Typography";

const theme = createTheme({
  palette: {
    secondary: {
      main: "#565F71",
    },
  },
});

const StyledButton = styled(Button)({
  color: "white",

  backgroundColor: "#565F71",
  height: "35px",
  padding: "5px 10px",
  borderRadius: "15px",
  boxShadow: "0px 4px 4px 0px rgba(0, 0, 0, 0.25)",
  transition: ".2s",

  "&:hover": {
    backgroundColor: "#565F71",
    boxShadow: "none",
    transform: "translate(0, 4px)",
  },
});

export function NewPostButton() {
  function clickHandler() {
    console.log("pushed new post button.");
  }

  return (
    <>
      <StyledButton onClick={clickHandler}>
        <Stack spacing='5px' direction='row' alignItems='center'>
          <EditOutlinedIcon sx={{ fontSize: "18px" }} />
          <Typography
            sx={{
              fontFamily: "Roboto",
              fontSize: "20px",
              fontWeight: 600,
            }}
          >
            新規投稿
          </Typography>
        </Stack>
      </StyledButton>
    </>
  );
}
