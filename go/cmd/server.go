// Copyright (C) 2018 Cranky Kernel
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/crankykernel/maker/go/server"
)

var ServerCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		server.ServerFlags.DataDirectory = DefaultDataDirectory
		server.ServerMain()
	},
}

func init() {
	flags := ServerCmd.Flags()
	flags.Int16VarP(&server.ServerFlags.Port, "port", "p", 6045, "Port")
	flags.StringVar(&server.ServerFlags.Host, "host", "127.0.0.1", "Host to bind to")
	flags.BoolVar(&server.ServerFlags.NoLog, "nolog", false, "Disable logging to file")
	flags.BoolVar(&server.ServerFlags.OpenBrowser, "open", false, "Open browser")
	rootCmd.AddCommand(ServerCmd)
}
