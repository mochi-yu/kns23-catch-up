import NotificationsNoneIcon from "@mui/icons-material/NotificationsNone";
import { IconButton } from "@mui/material";

export function NotifyButton() {
  function notifyButtonClickHandler() {}

  return (
    <>
      <IconButton sx={{ p: "0px" }} onClick={notifyButtonClickHandler}>
        <NotificationsNoneIcon sx={{ fontSize: "30px", color: "black" }} />
      </IconButton>
    </>
  );
}
