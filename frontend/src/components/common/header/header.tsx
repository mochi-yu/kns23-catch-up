import Box from "@mui/material/Box";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import NotificationsNoneIcon from "@mui/icons-material/NotificationsNone";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import { SearchBox } from "./search_box";
import { NewPostButton } from "@/components/new_post_button";

export function Header() {
  return (
    <>
      <Stack
        direction='row'
        sx={{ width: "100%", bgcolor: "#D6E3FF", py: "10px", px: "20px" }}
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
          <NotificationsNoneIcon sx={{ fontSize: "30px" }} />
          <AccountCircleIcon sx={{ fontSize: "50px" }} />
          <NewPostButton />
        </Stack>
      </Stack>
    </>
  );
}
