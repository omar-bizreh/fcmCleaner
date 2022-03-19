resource "kubernetes_namespace" "fcmcleaner" {
  metadata {
    name = "fcmcleaner-dev"
  }
}

resource "kubernetes_deployment" "fcmcleaner" {
  metadata {
    name      = "fcmcleaner-deployment"
    namespace = kubernetes_namespace.fcmcleaner.metadata.0.name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "fcmcleanerApi"
      }
    }
    template {
      metadata {
        labels = {
          app = "fcmcleanerApi"
        }
      }
      spec {
        container {
          image = "localhost:32000/fcmcleaner:r9"
          name  = "api-container"
          port {
            container_port = 80
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "fcmcleaner" {
  metadata {
    name      = "fcmcleaner-api-service"
    namespace = kubernetes_namespace.fcmcleaner.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment.fcmcleaner.spec.0.template.0.metadata.0.labels.app
    }
    type = "ClusterIP"
    port {
      port     = 80
      protocol = "TCP"
    }
  }
}

resource "kubernetes_ingress_v1" "fcmcleaner" {
  metadata {
    name = "fcmcleaner-ingress"
    annotations = {
      "kubernetes.io/ingress.class" = "nginx"
    }
    namespace = kubernetes_namespace.fcmcleaner.metadata.0.name
  }
  spec {
    rule {
      http {
        path {
          path = "/"
          backend {
            service {
              name = kubernetes_service.fcmcleaner.metadata.0.name
              port {
                number = 80
              }
            }
          }
        }
      }
    }
  }
}
