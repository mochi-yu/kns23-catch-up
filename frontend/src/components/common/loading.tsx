import { Backdrop, CircularProgress } from "@mui/material";

interface Props {
  open: boolean;
}

export const Loading = (props: Props) => {
  return (
    <Backdrop open={props.open} sx={{ zIndex: 20 }}>
      <CircularProgress color='inherit' />
    </Backdrop>
  );
};
