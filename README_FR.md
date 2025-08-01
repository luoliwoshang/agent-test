# Projet de démonstration GitHub Webhook

Un projet complet pour apprendre et démontrer les fonctionnalités des GitHub Webhooks, incluant une implémentation serveur Go et une documentation pédagogique détaillée.

## 🚀 Présentation du projet

Ce projet fournit un serveur de réception GitHub Webhook complet qui prend en charge :

- Réception et traitement de divers événements GitHub (Issues, Comments, Push, etc.)
- Vérification de signature HMAC-SHA256 pour assurer la sécurité
- Support pour le développement local et le déploiement sur la plateforme Render
- Sortie de logs détaillée et traitement d'événements
- Point de contrôle de santé

## 📁 Structure du projet

```
├── README.md              # Documentation du projet
├── webhook.md             # Tutoriel complet sur les Webhooks
├── claude-code-intro.md   # Introduction à l'outil Claude Code
├── go.mod                 # Configuration du module Go (répertoire racine)
├── render.yaml            # Configuration de déploiement Render
└── webhook-demo/          # Implémentation du serveur Webhook
    ├── go.mod            # Configuration du module Go
    └── server.go         # Code serveur principal
```

## 🛠️ Démarrage rapide

### Prérequis

- Go 1.21 ou version supérieure
- Git

### Développement local

1. **Cloner le projet**
   ```bash
   git clone <repository-url>
   cd agent-test
   ```

2. **Lancer le serveur**
   ```bash
   cd webhook-demo
   go run server.go
   ```

3. **Tester le serveur**
   ```bash
   curl http://localhost:8080/health
   ```

### Configuration du Webhook Secret (recommandé)

1. **Générer une clé secrète**
   ```bash
   openssl rand -hex 20
   ```

2. **Définir la variable d'environnement**
   ```bash
   export WEBHOOK_SECRET=your-generated-secret
   cd webhook-demo
   go run server.go
   ```

## 🌐 Déploiement sur Render

Le projet inclut un fichier de configuration `render.yaml` pour un déploiement en un clic sur la plateforme Render :

1. Connecter le dépôt GitHub à Render
2. Sélectionner Web Service
3. Render utilisera automatiquement la configuration de `render.yaml`

Après le déploiement, cela configurera automatiquement :
- La variable d'environnement `WEBHOOK_SECRET`
- Le contrôle de santé
- Le déploiement automatique

## 📚 Tutoriel d'utilisation

Pour un tutoriel d'utilisation détaillé, consultez :
- [**webhook.md**](webhook.md) - Tutoriel complet de configuration et d'utilisation des GitHub Webhooks
- Inclut des étapes détaillées pour la traversée NAT, la configuration GitHub, les méthodes de test, etc.

## 🎯 Événements pris en charge

Le serveur peut traiter les événements GitHub suivants :

- **ping** - Test de connexion Webhook
- **issues** - Création, édition, fermeture d'issues, etc.
- **issue_comment** - Commentaires d'issues (support de détection de commandes)
- **push** - Envoi de code
- D'autres événements seront également enregistrés

## 🔧 Points de terminaison API

- `POST /webhook` - Point de terminaison de réception d'événements Webhook
- `GET /health` - Point de terminaison de contrôle de santé

## 🔒 Fonctionnalités de sécurité

- Vérification de signature HMAC-SHA256
- Support du GitHub Webhook Secret
- Enregistrement détaillé des logs de sécurité
- Validation d'entrée et gestion d'erreurs

## 📖 Documentation connexe

- [Introduction à Claude Code](claude-code-intro.md) - Introduction à l'outil assistant de programmation IA
- [Tutoriel Webhook](webhook.md) - Enseignement complet depuis zéro

## 🤝 Contribution

Les issues et pull requests sont les bienvenues pour améliorer ce projet !

## 📄 Licence

Ce projet est uniquement destiné à des fins d'apprentissage et de démonstration.