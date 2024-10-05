import { Box, Toolbar, CssBaseline } from "@mui/material";
import { ThemeProvider } from "@mui/material/styles";
import Header from "../components/header";
import Sidebar from "../components/sidebar";
import { uecsTheme } from "../styles/theme";

export function Navigation({ children }: { children: React.ReactNode }) {
  return (
    <ThemeProvider theme={uecsTheme}>
      <Box sx={{ display: "flex" }}>
        <CssBaseline />
        <Header></Header>
        <Sidebar></Sidebar>
        <Box component="main" sx={{ flexGrow: 1, padding: 3 }}>
            <Toolbar />
            {children}
        </Box>
      </Box>
    </ThemeProvider>
  );
}
