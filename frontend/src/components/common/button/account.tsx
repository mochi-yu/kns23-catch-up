import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import { IconButton } from "@mui/material";

export function AccountButton() {
  function accountButtonClickHandler() {}

  return (
    <>
      <IconButton sx={{ p: "0px" }} onClick={accountButtonClickHandler}>
        <AccountCircleIcon sx={{ fontSize: "50px", color: "black" }} />
      </IconButton>
    </>
  );
}
