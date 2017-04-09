(setq-default c-basic-offset 4
			  tab-width 4
			  c-default-style "linux"
			  indent-tabs-mode t)
(require 'cc-mode)
    (add-to-list 'c-mode-common-hook
      (lambda () (setq c-syntactic-indentation nil)))
(provide 'init-c-mode)
