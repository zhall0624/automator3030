import { AppBar, Box, Typography, Button, Toolbar } from "@mui/material";
import { createRootRoute, Link, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";

export const Route = createRootRoute({
  component: () => (
    <>
      <Box>
        <AppBar position="static">
          <Toolbar>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              Automator3030
            </Typography>
            <Button component={Link} to="/incoming_webhooks" color="inherit">
              Incoming Webhooks
            </Button>
          </Toolbar>
        </AppBar>
        <Outlet />
        <TanStackRouterDevtools />
      </Box>
    </>
  ),
});
