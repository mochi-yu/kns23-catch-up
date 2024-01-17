import { Box, Typography } from "@mui/material";

interface Props {
  text: string
}

export function SearchBox(props: Props) {
  return (
    <>
      <Box width="100%" bgcolor="gray" height="100%">
        <Typography>{props.text}</Typography>
      </Box>
    </>
  )
}