package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var namespace string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List pods",
	Run: func(cmd *cobra.Command, args []string) {
		kubeConfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			log.Fatal(err)
		}

		clientSet, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatal(err)
		}

		list, err := clientSet.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, pod := range list.Items {
			fmt.Printf("Nome %s - Status %s\n", pod.Name, pod.Status.Phase)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Namespace")
}
