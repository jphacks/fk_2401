import {
  Box,
  Drawer,
  Link,
  List,
  ListItem,
  ListItemButton,
  Toolbar,
} from "@mui/material";
import HomeIcon from "@mui/icons-material/Home";
import SettingsRemoteIcon from "@mui/icons-material/SettingsRemote";
import TimelineIcon from "@mui/icons-material/Timeline";

const drawerWidth = 240;

interface SidebarItem {
  title: string;
  path: string;
  icon: JSX.Element;
}

const sidebarItems: SidebarItem[] = [
  {
    title: "ハウス状態",
    path: "/",
    icon: <HomeIcon />,
  },
  {
    title: "デバイス設定",
    path: "/devices",
    icon: <SettingsRemoteIcon />,
  },
  {
    title: "ワークフロー制御",
    path: "/workflow",
    icon: <TimelineIcon />,
  },
];

const SideBar = () => {
  return (
    <Drawer
      variant="permanent"
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        [`& .MuiDrawer-paper`]: {
          width: drawerWidth,
          boxSizing: "border-box",
          backgroundColor: "#E0E0E0",
        },
      }}
    >
      <Toolbar />
      <Box sx={{ overflow: "auto" }}>
        <List>
          {sidebarItems.map((item, index) => (
            <ListItem key={index} disablePadding>
              <ListItemButton sx={{ display: "flex", gap: 1 }}>
                {/* <ListItemIcon>
                  {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
                </ListItemIcon> */}
                {item.icon}
                <Link
                  href={item.path}
                  underline="none"
                  color="inherit"
                  sx={{
                    "&:hover": {
                      color: "#000",
                    },
                  }}
                >
                  {item.title}
                </Link>
              </ListItemButton>
            </ListItem>
          ))}
        </List>
      </Box>
    </Drawer>
  );
};

export default SideBar;
