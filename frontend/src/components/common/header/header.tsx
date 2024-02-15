"use client";

import Box from "@mui/material/Box";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import Slide from "@mui/material/Slide";
import AppBar from "@mui/material/AppBar";
import { useScrollTrigger } from "@mui/material";
import { SearchBox } from "./search_box";
import { NewPostButton } from "@/components/common/button/new_post";
import { AccountButton } from "@/components/common/button/account";
import { NotifyButton } from "@/components/common/button/notify";

interface Props {
  window?: () => Window;
}

export function Header(props: Props) {
  const window = props.window;
  const trigger = useScrollTrigger({
    target: window ? window() : undefined,
  });

  return (
    <>
      <Box sx={{ display: "flex" }} component='header'>
        <Slide appear={false} direction='down' in={!trigger}>
          <AppBar>
            <Stack
              component='header'
              direction='row'
              sx={{
                color: "black",

                width: "100%",
                bgcolor: "#D6E3FF",
                py: "10px",
                px: "20px",
              }}
              spacing='20px'
              alignItems='center'
            >
              <Typography
                style={{
                  fontFamily: "Roboto",
                  fontSize: "45px",
                  fontWeight: 600,
                }}
              >
                ReMeet
              </Typography>

              {/* 検索窓 */}
              <Box sx={{ flex: "1 1 0px" }} height='40px'>
                <SearchBox />
              </Box>

              <Stack spacing='5px' direction='row' alignItems='center'>
                <NotifyButton />
                <AccountButton />
                <NewPostButton />
              </Stack>
            </Stack>
          </AppBar>
        </Slide>
      </Box>
    </>
  );
}
