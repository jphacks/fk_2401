import { AppBar, Toolbar, Typography, Link } from "@mui/material";

export default function Header() {
  return (
    <AppBar
      position="fixed"
      sx={{
        zIndex: (theme) => theme.zIndex.drawer + 1,
        backgroundColor: "primary.main",
      }}
    >
      <Toolbar>
        <Typography variant="h6" noWrap component="div">
          <Link
            href="/"
            underline="none"
            sx={{
              color: "primary.contrastText",
              "&:hover": {
                color: "primary.contrastText",
              },
            }}
          >
            UECS Navi
          </Link>
        </Typography>
      </Toolbar>
    </AppBar>
  );
}
