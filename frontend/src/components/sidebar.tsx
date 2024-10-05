import {
  Box,
  Drawer,
  Link,
  List,
  ListItem,
  ListItemButton,
  Toolbar,
} from "@mui/material";

const drawerWidth = 240;

const SideBar = () => {
  const ListItems: string[] = [
    "ハウス状態",
    "デバイス設定",
  ]

  return (
    <Drawer
      variant="permanent"
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        [`& .MuiDrawer-paper`]: {
          width: drawerWidth,
          boxSizing: "border-box",
          backgroundColor: "primary.light",
        },
      }}
    >
      <Toolbar />
      <Box sx={{ overflow: "auto" }}>
        <List>
          {ListItems.map((item) => (
            <ListItem key={item} disablePadding>
              <ListItemButton>
                {/* <ListItemIcon>
                  {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
                </ListItemIcon> */}
                <Link
                  href="/" 
                  underline="none" 
                  color="inherit"
                  sx={{
                    color: "#000",
                    "&:hover": {
                      color: "#000", 
                    },
                  }}
                >
                  {item}
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
