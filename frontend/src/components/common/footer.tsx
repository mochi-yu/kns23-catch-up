"use client";

import Box from "@mui/material/Box";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import { Container } from "@mui/material";

export function Footer() {
  return (
    <>
      <Box>
        <Stack
          component='footer'
          direction='row'
          sx={{
            color: "white",
            width: "100%",
            bgcolor: "#2E3036",
            p: "10px",
          }}
        >
          <Container sx={{ textAlign: "center" }}>
            <Typography
              style={{
                fontFamily: "Roboto",
                fontSize: "12px",
                fontWeight: 500,
              }}
            >
              © 2024 神奈総23期有志
            </Typography>
          </Container>
        </Stack>
      </Box>
    </>
  );
}
