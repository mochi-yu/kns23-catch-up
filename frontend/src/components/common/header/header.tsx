import Box from "@mui/material/Box";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import NotificationsNoneIcon from '@mui/icons-material/NotificationsNone';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import EditIcon from '@mui/icons-material/Edit';
import { SearchBox } from "./searchBox";
import Button from "@mui/material/Button";

export function Header() {
  return (
    <>
      <Stack
        direction="row"
        sx={{width: "100%", bgcolor: "#D6E3FF", py: "10px", px: "20px"}}
        spacing="10px"
        alignItems="center"
      >
        <Typography variant="h4">
          ReMeet
        </Typography>
        
        {/* 検索窓 */}
        <Box sx={{flex: "1 1 0px"}} height="40px">
          <SearchBox text="aaa" />
        </Box>

        <NotificationsNoneIcon sx={{fontSize: "30px"}}/>
        <AccountCircleIcon sx={{fontSize: "50px"}}/>
        <Button sx={{bgcolor: "#565F71", color: "#FFFFFF", p: "5px", borderRadius: "15px", height: "40px"}}>
          <EditIcon />
          <Typography>
            新規投稿
          </Typography>
        </Button>
      </Stack>
    </>
  )
}